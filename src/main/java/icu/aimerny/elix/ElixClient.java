package icu.aimerny.elix;

import icu.aimerny.elix.entity.ModEntities;
import icu.aimerny.elix.entity.client.ModModelLayers;
import icu.aimerny.elix.entity.client.TigerModel;
import icu.aimerny.elix.entity.client.TigerRender;
import icu.aimerny.elix.registry.ModBlock;
import net.fabricmc.api.ClientModInitializer;
import net.fabricmc.fabric.api.blockrenderlayer.v1.BlockRenderLayerMap;
import net.fabricmc.fabric.api.client.rendering.v1.EntityModelLayerRegistry;
import net.fabricmc.fabric.api.client.rendering.v1.EntityRendererRegistry;
import net.minecraft.client.render.RenderLayer;
import net.minecraft.client.render.entity.EntityRenderer;
import net.minecraft.client.render.entity.model.EntityModelLayer;

public class ElixClient implements ClientModInitializer {
    @Override
    public void onInitializeClient() {
        BlockRenderLayerMap.INSTANCE.putBlock(ModBlock.OBSIDIAN_GLASS_BLOCK, RenderLayer.getTranslucent());
        EntityModelLayerRegistry.registerModelLayer(ModModelLayers.TIGER, TigerModel::getTexturedModelData);
        EntityRendererRegistry.register(ModEntities.TIGER, TigerRender::new);
    }
}
