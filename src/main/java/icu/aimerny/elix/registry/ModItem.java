package icu.aimerny.elix.registry;


import icu.aimerny.elix.Elix;
import icu.aimerny.elix.armor.ObsidianArmor;
import icu.aimerny.elix.consts.IdConst;
import icu.aimerny.elix.tool.ObsidianToolMaterial;
import net.fabricmc.fabric.api.item.v1.FabricItemSettings;
import net.minecraft.item.ArmorMaterial;
import net.minecraft.item.Item;
import net.minecraft.registry.Registries;
import net.minecraft.registry.Registry;
import net.minecraft.util.Identifier;

public class ModItem {

    public static final Item OBSIDIAN_INGOT = new Item(new FabricItemSettings());

    public static void register(String id, Item item) {
        Registry.register(Registries.ITEM, new Identifier(Elix.MOD_ID, id), item);
    }

    public static void init() {
        // init armors
        ObsidianArmor.init();
        ObsidianToolMaterial.init();

        // init other
        register(IdConst.OBSIDIAN_INGOT, OBSIDIAN_INGOT);
    }

}
