package icu.aimerny.elix.entity.client;

import icu.aimerny.elix.entity.custom.TigerEntity;
import net.minecraft.client.model.*;
import net.minecraft.client.render.VertexConsumer;
import net.minecraft.client.render.entity.model.EntityModel;
import net.minecraft.client.render.entity.model.SinglePartEntityModel;
import net.minecraft.client.util.math.MatrixStack;
import net.minecraft.entity.Entity;


// Made with Blockbench 4.9.4
// Exported for Minecraft version 1.17+ for Yarn
// Paste this class into your mod and generate all required imports
public class TigerModel<T extends TigerEntity> extends SinglePartEntityModel<T> {
	private final ModelPart bb_main;
	public TigerModel(ModelPart root) {
		this.bb_main = root.getChild("bb_main");
	}
	public static TexturedModelData getTexturedModelData() {
		ModelData modelData = new ModelData();
		ModelPartData modelPartData = modelData.getRoot();
		ModelPartData bb_main = modelPartData.addChild("bb_main", ModelPartBuilder.create().uv(0, 0).cuboid(-5.0F, -9.0F, -2.0F, 10.0F, 4.0F, 4.0F, new Dilation(0.0F))
		.uv(20, 8).cuboid(-4.0F, -7.0F, -2.0F, 2.0F, 7.0F, 1.0F, new Dilation(0.0F))
		.uv(6, 16).cuboid(-4.0F, -7.0F, 1.0F, 2.0F, 7.0F, 1.0F, new Dilation(0.0F))
		.uv(0, 16).cuboid(3.0F, -7.0F, 1.0F, 2.0F, 7.0F, 1.0F, new Dilation(0.0F))
		.uv(15, 15).cuboid(3.0F, -7.0F, -2.0F, 2.0F, 7.0F, 1.0F, new Dilation(0.0F))
		.uv(0, 8).cuboid(-8.0F, -12.0F, -2.0F, 4.0F, 4.0F, 4.0F, new Dilation(0.0F)), ModelTransform.pivot(0.0F, 24.0F, 0.0F));
		return TexturedModelData.of(modelData, 32, 32);
	}
	@Override
	public void setAngles(TigerEntity entity, float limbSwing, float limbSwingAmount, float ageInTicks, float netHeadYaw, float headPitch) {
	}
	@Override
	public void render(MatrixStack matrices, VertexConsumer vertexConsumer, int light, int overlay, float red, float green, float blue, float alpha) {
		bb_main.render(matrices, vertexConsumer, light, overlay, red, green, blue, alpha);
	}

	@Override
	public ModelPart getPart() {
		return this.bb_main;
	}
}